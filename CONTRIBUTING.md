# How to contribute

Performance is great deal to company, this library helps in that.
but ofcourse, nothing is perfect, so viewer might have ideas to change or propose.
Here we talk how to go ahead with that.

## Getting Started

* Make sure you have a [GitHub account](https://github.com/signup/free)
* Submit a ticket in [Compile Blades Project](https://github.com/Te-cho/compile-blades/projects/1) for your issue, assuming one does not already exist.
  * Clearly describe the issue including steps to reproduce when it is a bug.
  * Make sure you fill in the earliest version that you know has the issue.
* Fork the repository on GitHub

## Making Changes

* Create a topic branch from where you want to base your work.
  * This is usually the master branch.
  * Only target release branches if you are certain your fix must be on that
    branch.
  * To quickly create a topic branch based on master; `git checkout -b
    fix/master/my_contribution master`. Please avoid working directly on the
    `master` branch.
* Make commits of logical units.
* Check for unnecessary whitespace with `git diff --check` before committing.
* Make sure your commit messages are in the proper format.

````
    (PUP-1234) Make the example in CONTRIBUTING imperative and concrete

    Without this patch applied the example commit message in the CONTRIBUTING
    document is not a concrete example.  This is a problem because the
    contributor is left to imagine what the commit message should look like
    based on a description rather than an example.  This patch fixes the
    problem by making the example concrete and imperative.

    The first line is a real life imperative statement with a ticket number
    from our issue tracker.  The body describes the behavior without the patch,
    why this is a problem, and how the patch fixes the problem when applied.
````

* Make sure you have added the necessary tests for your changes.
* Run _all_ the tests to assure nothing else was accidentally broken.

## Making Trivial Changes

### Documentation

For changes of a trivial nature to comments and documentation, it is not
always necessary to create a new ticket in Jira. In this case, it is
appropriate to start the first line of a commit with '(doc)' instead of
a ticket number.

````
    (doc) Add documentation commit example to CONTRIBUTING

    There is no example for contributing a documentation commit
    to the Puppet repository. This is a problem because the contributor
    is left to assume how a commit of this nature may appear.

    The first line is a real life imperative statement with '(doc)' in
    place of what would have been the ticket number in a
    non-documentation related commit. The body describes the nature of
    the new documentation or comments added.
````

## Submitting Changes

* Push your changes to a topic branch in your fork of the repository.
* Submit a pull request to the repository in the Te-cho organization.
* Update your ticket to mark that you have submitted code and are ready for it to be reviewed (Status: Ready for Merge).
  * Include a link to the pull request in the ticket.
* The core team looks at Pull Requests on a regular basis in a weekly triage. 
* After feedback has been given we expect responses within two weeks. After two
  weeks we may close the pull request if it isn't showing any activity.

## Revert Policy
By running tests in advance and by engaging with peer review for prospective
changes, your contributions have a high probability of becoming long lived
parts of the the project. After being merged, the code will run through a
series of testing pipelines on a large number of operating system
environments. These pipelines can reveal incompatibilities that are difficult
to detect in advance.

If the code change results in a test failure, we will make our best effort to
correct the error. If a fix cannot be determined and committed within 24 hours
of its discovery, the commit(s) responsible _may_ be reverted, at the
discretion of the committer and Package maintainers. This action would be taken
to help maintain passing states in our testing pipelines.

The original contributor will be notified of the revert in the Github ticket
associated with the change. A reference to the test(s) and operating system(s)
that failed as a result of the code change will also be added to the Github
ticket. This test(s) should be used to check future submissions of the code to
ensure the issue has been resolved.

### Summary
* Changes resulting in test pipeline failures will be reverted if they cannot
  be resolved within one business day.

# Additional Resources

* [General GitHub documentation](https://help.github.com/)
* [GitHub pull request documentation](https://help.github.com/articles/creating-a-pull-request/)
