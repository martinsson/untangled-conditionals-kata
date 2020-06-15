package src

import "errors"

type Pipeline struct {
	config  Config
	emailer Emailer
	log     Logger
}

func (p *Pipeline) run(project Project) {

	var err error = nil

	if project.hasTests() {
		if "success" == project.runTests() {
			p.log.info("Tests passed")

		} else {
			err = errors.New("Tests failed")

		}
	} else {
		p.log.info("No tests")

	}

	if err == nil {
		if "success" == project.deploy() {
			p.log.info("Deployment successful")

		} else {
			err = errors.New("Deployment failed")
		}
	} else {
	}

	var endResult = ""
	if err != nil {
		p.log.error(err.Error())
		endResult = err.Error()
	} else {
		endResult = "Deployment completed successfully"
	}
	if p.config.sendEmailSummary() {
		p.log.info("Sending email")
		p.emailer.send(endResult)
	} else {
		p.log.info("Email disabled")
	}
}
