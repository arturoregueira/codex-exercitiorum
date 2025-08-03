package electionday

import (
	"fmt"
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {

	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {

	if counter != nil {
		return *counter
	} else {
		return 0
	}

}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	myElection := ElectionResult{
		Name:  candidateName,
		Votes: votes}

	return &myElection
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	myResult := *result
	name := myResult.Name
	count := myResult.Votes

	output := fmt.Sprintf("%s (%d)", name, count)

	return output

}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] = results[candidate] - 1

}
