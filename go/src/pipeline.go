package src

import "errors"

type Pipeline struct {
	config  Config
	emailer Emailer
	log     Logger
}

func (p *Pipeline) run(project Project) {

	// for stages
	err := p.runTestStage(project)
	if err == nil {
		err = p.runDeployStage(project)
	}

	// reporting
	if err != nil {
		p.log.error(err.Error())
	}

	result := p.computeEndResult(err)

	// email or null object
	if p.config.sendEmailSummary() {
		p.log.info("Sending email")
		p.emailer.send(result)
	} else {
		p.log.info("Email disabled")
	}
}

func (p *Pipeline) computeEndResult(err error) string {
	var endResult = "Deployment completed successfully"
	if err != nil {
		endResult = err.Error()
	}
	return endResult
}

func (p *Pipeline) runTestStage(project Project) error {
	err := error(nil)
	// almost possible to move this to project class, if it wasnt for the lack of logging in the failure case
	if project.hasTests() {
		if "success" == project.runTests() {
			p.log.info("Tests passed")
		} else {
			err = errors.New("Tests failed")
		}
	} else {
		p.log.info("No tests")
	}
	return err
}

func (p *Pipeline) runDeployStage(project Project) error {
	err := error(nil)
	if "success" == project.deploy() {
		p.log.info("Deployment successful")
	} else {
		err = errors.New("Deployment failed")
	}
	return err
}
