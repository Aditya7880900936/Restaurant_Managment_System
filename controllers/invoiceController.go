package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Aditya7880900936/Restaurant_Managment_System/database"
	"github.com/Aditya7880900936/Restaurant_Managment_System/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewerFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice items"})
		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice items"})
		}
		c.JSON(http.StatusOK, allInvoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching invoice item"})
			return
		}
		var invoiceViewer InvoiceViewerFormat
		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceViewer.Order_id = invoice.Order_id
		invoiceViewer.Payment_due_date = invoice.Payment_due_date
		invoiceViewer.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceViewer.Payment_method = *invoice.Payment_method
		}
		invoiceViewer.Invoice_id = invoice.Invoice_id
		invoiceViewer.Payment_status = *&invoice.Payment_status
		invoiceViewer.Payment_due = allOrderItems[0]["payment_due"]
		invoiceViewer.Table_number = allOrderItems[0]["table_number"]
		invoiceViewer.Order_details = allOrderItems[0]["order_items"]

		c.JSON(http.StatusOK, invoiceViewer)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
