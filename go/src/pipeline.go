package src

import "errors"

type Pipeline struct {
	config  Config
	emailer Emailer
	log     Logger
}
type pipelineStage func(project Project) (string, error)

func (p *Pipeline) run(project Project) {

	// for stages
	stages := []pipelineStage{testStage, deployStage}

	err := p.runStages(project, stages)

	result := p.computeEndResult(err)

	// email or null object
	if p.config.sendEmailSummary() {
		p.log.info("Sending email")
		p.emailer.send(result)
	} else {
		p.log.info("Email disabled")
	}
}

func (p *Pipeline) runStages(project Project, stages []pipelineStage) error {
	var err = error(nil)
	var message = ""
	for _, stage := range stages {
		message, err = stage(project)
		if err != nil {
			p.log.error(err.Error())
			break
		} else {
			p.log.info(message)
		}
	}
	return err
}

func (p *Pipeline) computeEndResult(err error) string {
	var endResult = "Deployment completed successfully"
	if err != nil {
		endResult = err.Error()
	}
	return endResult
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
