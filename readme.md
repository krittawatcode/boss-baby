# Boss Baby's Revenge

This is a Go project that provides a solution to the "Boss Baby's Revenge" problem. The project includes a set of Go files that implement the logic to solve the problem, as well as tests to ensure the solution's correctness.

## Problem Statement
Boss Baby, known for his cool and clever ways, deals with teasing from the neighborhood kids who shoot water guns at his house. In response, Boss Baby seeks revenge by shooting at least one shot back, but only if the kids have already shot at him first. As Boss Baby's assistant, your task is to check if Boss Baby has sought revenge for every shot aimed at him at least once and hasn't initiated any shooting himself.

## Input

A string (S, 1 <= len(S) <= 1,000,000) containing characters 'S' and 'R', where 'S' represents a shot and 'R' represents revenge. The string represents the sequence of events for one day.

## Output

Return "Good boy" if all shots have been revenged at least once and Boss Baby doesn’t initiate any shots to the neighborhood kids. Return "Bad boy" if these conditions are not satisfied.

### Example:
| Input    | Output   |
|----------|----------|
| SRSSRRR  | Good boy |
| RSSRR    | Bad boy  |
| SSSRRRRS | Bad boy  |
| SSRR     | Good boy |
| SRRSSR   | Bad boy  |

#### Explanation:
In the first example, the first shot “S” has been avenged by the second action. The second shots “SS” Have
been avenged by the following retaliation shots “RRR”.

In the second example, the first action is retaliation “R”, which makes Boss Baby a bad boy as he shouldn’t
shoot first.

In the third example, the first three shots “SSS” are revenged by at least 3 shots “RRRR”. However, the last
shot has no revenge, hence making Boss Baby a bad boy

### The main files in this project are:

- `bossBaby.go`: Contains the main logic for solving the problem.
- `bossBaby_test.go`: Contains the tests for the `bossBaby.go` file.
- `main.go`: The entry point of the application.

The project uses Go modules for dependency management, as specified in the [`go.mod`](command:_github.copilot.openRelativePath?%5B%22go.mod%22%5D) file.
