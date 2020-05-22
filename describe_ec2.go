package ec2
import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)


var (
	client        *ec2.EC2
	sortKey       []string
	latestAmiID   string
	previousAmiID string
)

func init() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("******")},
		Profile: "******",
	})
	if err != nil {
		fmt.Println("Error creating session ", err.Error())
		return

	}
	client = ec2.New(sess)

}

func GetLatestandPreviousAmiID() {
	resp, _ := client.DescribeImages(&ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("name"),
        //Change Value
				Values: []*string{aws.String("***********")},
			},
		}})

	for idx := range resp.Images {
		sortKey = append(sortKey, *resp.Images[idx].CreationDate)
	}
	sort.Strings(sortKey)
	for idx := range resp.Images {
		if *resp.Images[idx].CreationDate == sortKey[len(sortKey)-1] {
			latestAmiID = *resp.Images[idx].ImageId
			fmt.Println("Latest AMI-ID--> ", latestAmiID, *resp.Images[idx].CreationDate)
			break
		}
	}
 }
