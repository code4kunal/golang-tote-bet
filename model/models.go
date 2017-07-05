package model
/**
  *  This package consists of the data model
  *  used to consume the data within the task.
  */


type WinPoolObject struct {
	HorseID string
	Stake   string
}

type PlacePoolObject struct {
	HorseID string
	Stake   string
}

type ExactaPoolObject struct {
	HorseID string
	Stake   string
}

type QuinellaPoolObject struct {
	HorseID string
	Stake   string
}


type ResultObject struct {
	Result []string
}

