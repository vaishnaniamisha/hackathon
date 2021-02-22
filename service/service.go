package service

import "scripbox/hackathon/lib/database"

//HackathonService strucure
type ChallengeService struct {
	DbClient database.DBClientInterface
}
