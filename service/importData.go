package service

import (
	"furst/model"
	"furst/repository"

	"bytes"
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Payload struct{}

type ImportDataService struct{}

func (ImportDataService) Call(year int, month int) (int64, error) {
	status, err := DownloadCSV()
	if err != nil {
		return status, err
	}

	err = ImportFromCSV(year, month)
	if err != nil {
		return status, err
	}

	return status, nil
}

func DownloadCSV() (int64, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return 1, err
	}

	region := os.Getenv("AWS_REGION")
	access_key := os.Getenv("AWS_ACCESS_KEY")
	secret_key := os.Getenv("AWS_SECRET_KEY")
	arn := os.Getenv("AWS_LAMBDA_ARN")

	payload := Payload{}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return 1, err
	}

	client := lambda.New(session.New(), &aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     access_key,
			SecretAccessKey: secret_key,
		}),
	})

	params := &lambda.InvokeInput{
		FunctionName: aws.String(arn),
		Payload:      jsonBytes,
	}

	res, err := client.Invoke(params)
	if err != nil {
		return *res.StatusCode, err
	}

	return *res.StatusCode, nil
}

func ImportFromCSV(year int, month int) error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}

	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUDGET")
	key := "csv/" + strconv.Itoa(year) + "-" + strconv.Itoa(month) + ".csv"
	access_key := os.Getenv("AWS_ACCESS_KEY")
	secret_key := os.Getenv("AWS_SECRET_KEY")

	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     access_key,
			SecretAccessKey: secret_key,
		}),
	})

	obj, _ := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	defer obj.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	reader := transform.NewReader(strings.NewReader(buf.String()), japanese.ShiftJIS.NewDecoder())
	csvReader := csv.NewReader(reader)

	rows, err := csvReader.ReadAll()
	rows = rows[1:]

	for _, row := range rows {
		mf := row[9]
		payRepo := repository.PayRepository{}
		found := payRepo.FindByMf(mf)
		if found {
			continue
		}

		amount, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			panic(err)
		}
		dateStr := strings.Replace(row[1], "/", "-", -1)
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			panic(err)
		}
		content := row[2]
		memo := row[7]
		countIn, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}
		trans, err := strconv.Atoi(row[8])
		if err != nil {
			panic(err)
		}
		// TODO: for now.
		userID := uint(1)
		walletName := row[4]
		mainCategoryName := row[5]
		subCategoryName := row[6]

		walletRepo := repository.WalletRepository{}
		wallet, found := walletRepo.FindByName(walletName)
		if !found {
			wallet = model.Wallet{
				Name:   walletName,
				Amount: 0,
				UserID: userID,
			}
			err = walletRepo.SetWallet(&wallet)
		}
		if err != nil {
			panic(err)
		}

		mainRepo := repository.MainCategoryRepository{}
		mainCategory, found := mainRepo.FindByName(mainCategoryName)
		if !found {
			mainCategory = model.MainCategory{
				Name:   mainCategoryName,
				UserID: userID,
			}
			err = mainRepo.SetMainCategory(&mainCategory)
		}
		if err != nil {
			panic(err)
		}

		subRepo := repository.SubCategoryRepository{}
		subCategory, found := subRepo.FindByName(subCategoryName)
		if !found {
			subCategory = model.SubCategory{
				Name:           subCategoryName,
				UserID:         userID,
				MainCategoryID: mainCategory.ID,
			}
			err = subRepo.SetSubCategory(&subCategory)
		}
		if err != nil {
			panic(err)
		}

		pay := model.Pay{
			Amount:         -amount,
			Date:           date,
			Content:        content,
			Memo:           memo,
			CountIn:        (countIn == 1),
			Transfer:       (trans == 1),
			Mf:             mf,
			UserID:         userID,
			WalletID:       wallet.ID,
			MainCategoryID: mainCategory.ID,
			SubCategoryID:  subCategory.ID,
		}
		err = payRepo.SetPay(&pay)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
