package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"bookstore/pkg/utils"
	"bookstore/pkg/models"
)

var NewBook models.Book