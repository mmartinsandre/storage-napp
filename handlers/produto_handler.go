package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

type Produto struct {
    gorm.Model
    Codigo           string  `json:"codigo"`
    Nome             string  `json:"nome"`
    EstoqueTotal     int     `json:"estoque_total"`
    EstoqueCorte     int     `json:"estoque_corte"`
    EstoqueDisponivel int    `json:"estoque_disponivel"`
    PrecoDe          float64 `json:"preco_de"`
    PrecoPor         float64 `json:"preco_por"`
}

func getDB(c *gin.Context) *gorm.DB {
    return c.MustGet("db").(*gorm.DB)
}

func CreateProduto(c *gin.Context) {
    db := getDB(c)
    var produto Produto
    if err := c.BindJSON(&produto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    produto.EstoqueDisponivel = produto.EstoqueTotal - produto.EstoqueCorte

    if produto.PrecoDe < produto.PrecoPor {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Preço de não pode ser inferior ao preço por"})
        return
    }

    db.Create(&produto)
    c.JSON(http.StatusCreated, produto)
}

func GetProduto(c *gin.Context) {
    db := getDB(c)
    var produto Produto
    id := c.Param("id")
    if err := db.First(&produto, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
        return
    }
    c.JSON(http.StatusOK, produto)
}

func GetAllProdutos(c *gin.Context) {
    db := getDB(c)
    var produtos []Produto
    if err := db.Find(&produtos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar produtos"})
        return
    }
    c.JSON(http.StatusOK, produtos)
}

func UpdateProduto(c *gin.Context) {
    db := getDB(c)
    var produto Produto
    id := c.Param("id")
    if err := db.First(&produto, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
        return
    }

    if err := c.BindJSON(&produto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    produto.EstoqueDisponivel = produto.EstoqueTotal - produto.EstoqueCorte

    if produto.PrecoDe < produto.PrecoPor {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Preço de não pode ser inferior ao preço por"})
        return
    }

    db.Save(&produto)
    c.JSON(http.StatusOK, produto)
}

func DeleteProduto(c *gin.Context) {
    db := getDB(c)
    var produto Produto
    id := c.Param("id")
    if err := db.First(&produto, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
        return
    }

    db.Delete(&produto)
    c.JSON(http.StatusNoContent, gin.H{})
}
