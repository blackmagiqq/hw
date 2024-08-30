package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestListSuite struct {
	suite.Suite
	list *list
}

func (suite *TestListSuite) SetupTest() {
	listItem1 := &ListItem{Value: 1}
	listItem2 := &ListItem{Value: 2}
	listItem1.Next = listItem2
	listItem2.Prev = listItem1

	suite.list = &list{
		countOfListItems: 2,
		frontItem:        listItem1,
		backItem:         listItem2,
	}
}

func (suite *TestListSuite) TestPushFrontPositive() {
	expectedFrontListItem := &ListItem{
		Value: 3,
		Next:  suite.list.frontItem,
	}
	expectedBackListItem := suite.list.backItem
	expectedCountOfListItems := 3

	suite.list.PushFront(3)

	require.Equal(suite.T(), expectedCountOfListItems, suite.list.countOfListItems, "unexpected count of list items")
	require.Equal(suite.T(), expectedFrontListItem, suite.list.frontItem, "unexpected front ListItem")
	require.Equal(suite.T(), expectedBackListItem, suite.list.backItem, "unexpected back ListItem")
}

func (suite *TestListSuite) TestPushBackPositive() {
	expectedFrontListItem := suite.list.frontItem
	expectedBackListItem := &ListItem{
		Value: 3,
		Prev:  suite.list.backItem,
	}

	expectedCountOfListItems := 3

	suite.list.PushBack(3)

	require.Equal(suite.T(), expectedCountOfListItems, suite.list.countOfListItems, "unexpected count of list items")
	require.Equal(suite.T(), expectedFrontListItem, suite.list.frontItem, "unexpected front ListItem")
	require.Equal(suite.T(), expectedBackListItem, suite.list.backItem, "unexpected back ListItem")
}

func (suite *TestListSuite) TestRemovePositive() {
	suite.list.PushBack(3)

	expectedFrontListItem := suite.list.frontItem
	expectedBackListItem := &ListItem{
		Value: 3,
		Prev:  expectedFrontListItem,
	}

	expectedCountOfListItems := 2

	suite.list.Remove(suite.list.frontItem.Next)

	require.Equal(suite.T(), expectedCountOfListItems, suite.list.countOfListItems, "unexpected count of list items")
	require.Equal(suite.T(), expectedFrontListItem, suite.list.frontItem, "unexpected front ListItem")
	require.Equal(suite.T(), expectedBackListItem, suite.list.backItem, "unexpected back ListItem")
}

func (suite *TestListSuite) TestMoveToFrontPositive() {
	newListItem := suite.list.PushBack(3)

	expectedFrontListItem := &ListItem{
		Value: 3,
		Next:  suite.list.frontItem,
	}
	expectedBackListItem := suite.list.frontItem.Next

	expectedCountOfListItems := 3

	suite.list.MoveToFront(newListItem)

	require.Equal(suite.T(), expectedCountOfListItems, suite.list.countOfListItems, "unexpected count of list items")
	require.Equal(suite.T(), expectedFrontListItem, suite.list.frontItem, "unexpected front ListItem")
	require.Equal(suite.T(), expectedBackListItem, suite.list.backItem, "unexpected back ListItem")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestListSuite))
}
