package reponosql

import (
	"strconv"
	"sub-hf-product-p5/internal/core/db"
	"sub-hf-product-p5/internal/core/domain/entity/dto"
	"sub-hf-product-p5/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var _ repository.ProductRepository = (*productDB)(nil)

type productDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewProductRepository(database db.NoSQLDatabase, tableName string) *productDB {
	return &productDB{Database: database, tableName: tableName}
}

func (p *productDB) GetProductByID(uuid string) (*dto.ProductDB, error) {
	partitionKeyName := "uuid"

	input := &dynamodb.QueryInput{
		TableName:              aws.String(p.tableName),
		KeyConditionExpression: aws.String("#pk = :value"),
		ExpressionAttributeNames: map[string]string{
			"#pk": partitionKeyName,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: uuid},
		},
	}

	result, err := p.Database.Query(input)
	if err != nil {
		return nil, err
	}

	var productList = make([]dto.ProductDB, 0)
	for _, item := range result.Items {
		var pDb dto.ProductDB
		if err := attributevalue.UnmarshalMap(item, &pDb); err != nil {
			return nil, err
		}
		productList = append(productList, pDb)
	}

	if len(productList) > 0 {
		return &productList[0], nil
	}

	return nil, nil
}

func (p *productDB) SaveProduct(product dto.ProductDB) (*dto.ProductDB, error) {

	priceString := strconv.FormatFloat(product.Price, 'f', -1, 64)
	putItem := map[string]types.AttributeValue{
		"uuid": &types.AttributeValueMemberS{
			Value: product.UUID,
		},
		"name": &types.AttributeValueMemberS{
			Value: product.Name,
		},
		"category": &types.AttributeValueMemberS{
			Value: product.Category,
		},
		"image": &types.AttributeValueMemberS{
			Value: product.Image,
		},
		"description": &types.AttributeValueMemberS{
			Value: product.Description,
		},
		"price": &types.AttributeValueMemberN{
			Value: priceString,
		},
		"createdAt": &types.AttributeValueMemberS{
			Value: product.CreatedAt,
		},
		"deactivatedAt": &types.AttributeValueMemberS{
			Value: product.DeactivatedAt,
		},
	}

	inputPutItem := &dynamodb.PutItemInput{
		Item:      putItem,
		TableName: aws.String(p.tableName),
	}

	putOut, err := p.Database.PutItem(inputPutItem)

	if err != nil {
		return nil, err
	}

	var out dto.ProductDB

	if err := attributevalue.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	out = dto.ProductDB{
		UUID:          product.UUID,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	return &out, nil
}

func (p *productDB) UpdateProductByID(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {

	update := expression.Set(expression.Name("name"), expression.Value(product.Name))
	update.Set(expression.Name("category"), expression.Value(product.Category))
	update.Set(expression.Name("image"), expression.Value(product.Image))
	update.Set(expression.Name("description"), expression.Value(product.Description))
	update.Set(expression.Name("price"), expression.Value(product.Price))
	update.Set(expression.Name("createdAt"), expression.Value(product.CreatedAt))
	update.Set(expression.Name("deactivatedAt"), expression.Value(product.DeactivatedAt))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	inputUpdateItem := &dynamodb.UpdateItemInput{
		TableName: aws.String(p.tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{
				Value: uuid,
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	updateOut, err := p.Database.UpdateItem(inputUpdateItem)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var out dto.ProductDB

	if err := attributevalue.UnmarshalMap(updateOut.Attributes, &out); err != nil {
		return nil, err
	}

	out = dto.ProductDB{
		UUID:          uuid,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	return &out, nil
}

func (p *productDB) GetProductByCategory(category string) ([]dto.ProductDB, error) {
	filter := "category = :value"
	attrSearch := map[string]types.AttributeValue{
		":value": &types.AttributeValueMemberS{
			Value: category,
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(p.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := p.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var productList = make([]dto.ProductDB, 0)
	for _, item := range result.Items {
		var pDb dto.ProductDB
		if err := attributevalue.UnmarshalMap(item, &pDb); err != nil {
			return nil, err
		}
		productList = append(productList, pDb)
	}

	if len(productList) > 0 {
		return productList, nil
	}

	return nil, nil
}

func (p *productDB) DeleteProductByID(uuid string) error {
	key := map[string]types.AttributeValue{
		"uuid": &types.AttributeValueMemberS{
			Value: uuid,
		},
	}

	inputUpdateItem := &dynamodb.DeleteItemInput{
		TableName: aws.String(p.tableName),
		Key:       key,
	}

	if _, err := p.Database.DeleteItem(inputUpdateItem); err != nil {
		return err
	}

	return nil
}
