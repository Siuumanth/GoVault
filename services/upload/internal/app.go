package internal

// import (
// 	"context"
// 	//core "govault/services/upload/internal/core"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// type App struct {
// 	Storage core.Storage
// 	// other deps like DAO, router...
// }

// func NewApp() (*App, error) {
// 	ctx := context.Background()

// 	// 1. load AWS config
// 	awsCfg, err := config.LoadDefaultConfig(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 2. create S3 client
// 	s3Client := s3.NewFromConfig(awsCfg)

// 	// 3. nject into your storage implementation
// 	storage := &core.S3Storage{
// 		Client: s3Client,
// 		Bucket: "govault-files",
// 	}

// 	return &App{
// 		Storage: storage,
// 	}, nil
// }
