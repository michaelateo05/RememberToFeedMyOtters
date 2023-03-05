package main

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)