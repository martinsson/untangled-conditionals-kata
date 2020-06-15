package src


type Pipeline struct {
    config Config
    emailer Emailer
    log Logger
}

func (p *Pipeline) run(project Project) {
    var testsPassed = false

	var endResult = ""

    if project.hasTests() {
        if "success" == project.runTests() {
            p.log.info("Tests passed")
            testsPassed = true
        } else {
            p.log.error("Tests failed")
            endResult = "Tests failed"
            testsPassed = false
        }
    } else {
        p.log.info("No tests")
        testsPassed = true
    }

    if testsPassed {
        if "success" == project.deploy() {
            p.log.info("Deployment successful")
            endResult = "Deployment completed successfully"

		} else {
            p.log.error("Deployment failed")
            endResult = "Deployment failed"

		}
    } else {

	}

    if p.config.sendEmailSummary() {
        p.log.info("Sending email")
        p.emailer.send(endResult)
    } else {
        p.log.info("Email disabled")
    }
}

