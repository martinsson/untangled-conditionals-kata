package src

import "errors"

type Project struct {
	buildsSuccessfully bool
	testStatus         TestStatus
}

func (p *Project) SetTestStatus(testStatus TestStatus) {
	p.testStatus = testStatus
}

func (p Project) hasTests() bool {
	return p.testStatus != NO_TESTS
}

func (p Project) runTests() string {
	if p.testStatus == PASSING_TESTS {
		return "success"
	}
	return "failure"
}
func (p Project) deploy() string {
	if p.buildsSuccessfully {
		return "success"
	}
	return "failure"
}

type ProjectBuilder struct {
	buildsSuccessfully bool
	testStatus         TestStatus
}

func builder() ProjectBuilder {
	return ProjectBuilder{}
}

func testStage(project Project) (string, error) {
	err := error(nil)
	var message = ""
	if project.hasTests() {
		if "success" == project.runTests() {
			message = "Tests passed"
		} else {
			err = errors.New("Tests failed")
		}
	} else {
		message = "No tests"
	}

	return message, err
}

func deployStage(project Project) (string, error) {
	err := error(nil)
	var message = ""
	if "success" == project.deploy() {
		message = "Deployment successful"

	} else {
		err = errors.New("Deployment failed")
	}
	return message, err
}
