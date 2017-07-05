package helper
/**
  * This package consists of the helper functions
  * to compute the objective within the task.
  */


import (
	"strings"
	"errors"
	"io/ioutil"
	"log"
	"github.com/Sirupsen/logrus"
	"github.com/walkover-task/model"
	"strconv"
)
/**
  * Parses and Validates result string
  * @param str
  * @return []string, error
*/
func ValidateResultString(str string) ( []string , error ) {
	s := strings.Split(str, ":")
	if len(s) < 5  {
		return nil, errors.New("Invalid Result String")
	}
	return s , nil
}

/**
  * Parses input stakes and populate all
  * the required pools.
  * @param filename
  * @return WinPool, PlacePool, ExactaPool, QuintellaPool
*/

func ParseInputAndPopulatePools (fileName string) ([]model.WinPoolObject, []model.PlacePoolObject,
	[]model.ExactaPoolObject, []model.QuinellaPoolObject, error){

	winPool := []model.WinPoolObject{}
	placePool := []model.PlacePoolObject{}
	exactaPool := []model.ExactaPoolObject{}
	quinellaPool := []model.QuinellaPoolObject{}


	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil,nil,nil,nil,err
	}

	lines := strings.Split(string(content), "\n")
	for _, line :=  range lines{
		parsedStakeList, errs := ParseStakes(line)
		if err!=nil{
			logrus.Error(errs)
			return nil,nil,nil,nil, err
		}

		if parsedStakeList[0]== "W"{
			winStakeObject := &model.WinPoolObject{
				HorseID:parsedStakeList[1],
				Stake:parsedStakeList[2],
			}
			winPool = append(winPool,*winStakeObject)
		}else if parsedStakeList[0]== "P"{
			PlaceStakeObject := &model.PlacePoolObject{
				HorseID:parsedStakeList[1],
				Stake:parsedStakeList[2],
			}
			placePool = append(placePool,*PlaceStakeObject)
		}else if parsedStakeList[0]== "E"{
			ExactaStakeObject := &model.ExactaPoolObject{
				HorseID:parsedStakeList[1],
				Stake:parsedStakeList[2],
			}
			exactaPool = append(exactaPool,*ExactaStakeObject)
		}else if parsedStakeList[0]== "Q"{
			QuinellaStakeObject := &model.QuinellaPoolObject{
				HorseID:parsedStakeList[1],
				Stake:parsedStakeList[2],
			}
			quinellaPool = append(quinellaPool,*QuinellaStakeObject)
		}else {
			err =  errors.New("Invalid Stake String")
			return nil, nil, nil, nil, err
		}


	}

	return winPool, placePool,exactaPool,quinellaPool,nil
}


/**
 * Parses and Validates single stake entry
 * @param stakeStr
 * @return []string, error
*/
func ParseStakes(stakesStr string)([]string, error){
	s := strings.Split(stakesStr, ":")
	if len(s) < 3 || len(s) == 0 {
		return nil, errors.New("Invalid Stake")
	}
	return s , nil

}



/**
  * Calculate dividends for each pool
  * @param str
  * @param totalStake
  * @param totalWinnerStake
  * @param res
  * @return error
*/
func CalculateDividends(pool string, totalStake float64, totalWinnerStake float64, res string) (error){

	if pool == "win"{
		dividend := ((totalStake * (0.85))/ totalWinnerStake)
		log.Printf("Win- Runner %s -%.2f",res,dividend)
	}else if pool == "place"{
		dividend := ((totalStake * (0.88))/ totalWinnerStake)
		log.Printf("Place- Runner %s -%.2f",res,dividend)
	}else if pool == "exacta"{
		dividend := ((totalStake * (0.82))/ totalWinnerStake)
		log.Printf("Place- Runners %s -%.2f",res,dividend)
	}else if pool == "quintella"{
		dividend := ((totalStake * (0.82))/ totalWinnerStake)
		log.Printf("Place- Runners %s -%.2f",res,dividend)
	}

	return nil

}


/**
  * Calculate meta data for Winpool
  * @param winPool
  * @param res
  * @return error
*/

func CalculateMetaForWinPool(win []model.WinPoolObject,res model.ResultObject) (error){

   totalPoolStake := 0.00
   totalWinnerStake := 0.00

	for _, wObj := range win{
		stake ,_ :=strconv.Atoi(wObj.Stake)
	   totalPoolStake = totalPoolStake + float64(stake)
		if wObj.HorseID == res.Result[1]{
			totalWinnerStake = totalWinnerStake + float64(stake)
		}

	}

	err := CalculateDividends("win",totalPoolStake,totalWinnerStake, res.Result[1])
	if err != nil {
		return err
	}
  return  nil
}


/**
  * Calculate meta data for PlacePool
  * @param placePool
  * @param res
  * @return error
*/

func CalculateMetaForPlacePool(place []model.PlacePoolObject,res model.ResultObject) (error){

	totalPoolStake := 0.00
	totalWinnerStakeFirstPosition := 0.00
	totalWinnerStakeSecondPosition := 0.00
	totalWinnerStakeThirdPosition := 0.00

	for _, pObj := range place{
		stake ,_ :=strconv.Atoi(pObj.Stake)
		totalPoolStake = totalPoolStake + float64(stake)
		if pObj.HorseID == res.Result[1]{
			totalWinnerStakeFirstPosition = totalWinnerStakeFirstPosition + float64(stake)
		}else if pObj.HorseID == res.Result[2]{
			totalWinnerStakeSecondPosition = totalWinnerStakeSecondPosition + float64(stake)
		}else if pObj.HorseID == res.Result[3]{
			totalWinnerStakeThirdPosition = totalWinnerStakeThirdPosition + float64(stake)
		}

	}

	err := CalculateDividends("place",float64(totalPoolStake/3.00),totalWinnerStakeFirstPosition, res.Result[1])
	if err != nil {
		return err
	}

	err = CalculateDividends("place",float64(totalPoolStake/3.00),totalWinnerStakeSecondPosition, res.Result[2])
	if err != nil {
		return err
	}

	err = CalculateDividends("place",float64(totalPoolStake/3.00),totalWinnerStakeThirdPosition, res.Result[3])
	if err != nil {
		return err
	}

	return nil
}

/**
  * Calculate meta data for ExactaPool
  * @param exactaPool
  * @param res
  * @return error
*/

func CalculateMetaForExactaPool(exacta []model.ExactaPoolObject,res model.ResultObject) (error){

	totalPoolStake := 0.00
	totalWinnerStake := 0.00
	resString := res.Result[1]+","+res.Result[2]

	for _, eObj := range exacta{
		stake ,_ :=strconv.Atoi(eObj.Stake)
		totalPoolStake = totalPoolStake + float64(stake)
		Id := (strings.Split(eObj.HorseID,","))
		if Id[0] == res.Result[1] && Id[1] == res.Result[2]{
			totalWinnerStake = totalWinnerStake + float64(stake)
		}

	}

	err := CalculateDividends("exacta",totalPoolStake,totalWinnerStake, resString)
	if err != nil {
		return err
	}
	return  nil
}


/**
  * Calculate meta data for QuinettaPool
  * @param quinettaPool
  * @param res
  * @return error
*/

func CalculateMetaForQuintellaPool(quintella []model.QuinellaPoolObject,res model.ResultObject) (error){

	totalPoolStake := 0.00
	totalWinnerStake := 0.00
	resString := res.Result[1]+","+res.Result[2]

	for _, eObj := range quintella{
		stake ,_ :=strconv.Atoi(eObj.Stake)
		totalPoolStake = totalPoolStake + float64(stake)
		Id := (strings.Split(eObj.HorseID,","))
		if Id[0] == res.Result[1] && Id[1] == res.Result[2]{
			totalWinnerStake = totalWinnerStake + float64(stake)
		}

	}

	err := CalculateDividends("quintella",totalPoolStake,totalWinnerStake, resString)
	if err != nil {
		return err
	}
	return  nil
}