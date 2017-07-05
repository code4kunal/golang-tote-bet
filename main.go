package main

import (
	"flag"
	"github.com/walkover-task/helper"
	"github.com/Sirupsen/logrus"
	"github.com/walkover-task/model"



)


func main() {


	resultStr := flag.String("result", "null", "a string representing result")

	flag.Parse()
	resVarString := *resultStr
	res, err := helper.ValidateResultString(resVarString)
	if err!=nil{
		logrus.Error(err)
		return
	}

	resultObject := model.ResultObject{}
    resultObject.Result =res

	winpool, placepool, exactapool, quintellapool, err := helper.ParseInputAndPopulatePools("stakes.txt")
	if err!=nil{
		logrus.Error(err)
		return
	}

	err = helper.CalculateMetaForWinPool(winpool,resultObject)
	if err!=nil{
		logrus.Error(err)
		return
	}

	err = helper.CalculateMetaForPlacePool(placepool,resultObject)
	if err!=nil{
		logrus.Error(err)
		return
	}

	err = helper.CalculateMetaForExactaPool(exactapool,resultObject)
	if err!=nil{
		logrus.Error(err)
		return
	}

	err = helper.CalculateMetaForQuintellaPool(quintellapool,resultObject)
	if err!=nil{
		logrus.Error(err)
		return
	}
}
