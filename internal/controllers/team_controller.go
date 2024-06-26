package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjayheaven/ggb/internal/models"
	"github.com/sanjayheaven/ggb/internal/pkg/firebase"
)

type CreateTeamRequest struct {
	TeamName   string `json:"teamName" binding:"required"`
	TeamLeader string `json:"teamLeader" binding:"required,email"`
}
func CreateTeam(c *gin.Context) {
    var req CreateTeamRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    team := models.Team{
        TeamName:    req.TeamName,
        TeamLeader:  req.TeamLeader,
        TeamMembers: []string{req.TeamLeader},
        TeamLength:  1,
        TeamCode:    models.GenerateTeamCode(),
    }

	ctx := context.Background()
	// Check if Team Name is taken
	iter := firebase.FirestoreClient.Collection("teams").Where("TeamName", "==", req.TeamName).Documents(ctx)
	defer iter.Stop()
	if _, err := iter.Next(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Team name already in use"})
		return
	}
	// Check if Team Leader is leader already
	iter = firebase.FirestoreClient.Collection("teams").Where("TeamLeader", "==", req.TeamLeader).Documents(ctx)
	defer iter.Stop()
	if _, err := iter.Next(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Team leader already in use"})
		return
	}
	// Add the Team 
    _, _, err := firebase.FirestoreClient.Collection("teams").Add(c, team)
    if err != nil {
        log.Printf("Error creating team: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
        return
    }
	
    c.JSON(http.StatusCreated, gin.H{
        "message":  "Team created successfully",
        "teamCode": team.TeamCode,
    })
}