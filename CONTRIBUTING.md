# Contributing to Goava

We'd love for you to contribute to our source code and to make Goava even better
than it is today!
Here are the guidelines we'd like you to follow:

 - [Code of Conduct](#coc)
 - [Question or Problem?](#question)
 - [Issues and Bugs](#issue)
 - [Feature Requests](#feature)
 - [Submission Guidelines](#submit)

## <a name="coc"></a> Code of Conduct

As contributors and maintainers of the Goava project, we pledge to respect
everyone who contributes by posting issues, updating documentation, submitting
pull requests, providing feedback in comments, and any other activities.

Communication through any of Goava's channels (GitHub, etc.) must be
constructive and never resort to personal attacks, trolling, public or private
harassment, insults, or other unprofessional conduct.

We promise to extend courtesy and respect to everyone involved in this project
regardless of gender, gender identity, sexual orientation, disability, age,
race, ethnicity, religion, or level of experience.
We expect anyone contributing to the Goava project to do the same.

If any member of the community violates this code of conduct, the maintainers of
the Goava project may take action, removing issues, comments, and PRs as deemed
appropriate.

If you are subject to or witness unacceptable behavior, or have any other
concerns, please drop us a line at gschauer.abc.inc+goava@gmail.com.

## <a name="question"></a> Got a Question or Problem?

If you have questions about how to use Goava, please direct these to
[StackOverflow][stackoverflow] and use the `goava` tag.
We are also available on GitHub issues.

If you feel that we're missing an important bit of documentation, feel free to
file an issue so we can help.

## <a name="issue"></a> Found an Issue?
If you find a bug in the source code or a mistake in the documentation, you can
help us by submitting an issue to our [GitHub Repository][github].
Even better you can submit a Pull Request with a fix.

See [below](#submit) for some guidelines.

## <a name="feature"></a> Want a Feature?
You can request a new feature by submitting an issue to our
[GitHub Repository][github].

If you would like to implement a new feature then consider what kind of change
it is:

* **Major Changes** that you wish to contribute to the project should be
discussed first on our [issue tracker][] so that we can better coordinate our
efforts, prevent duplication of work, and help you to craft the change so that
it is successfully accepted into the project.
* **Small Changes** can be crafted and submitted to the
[GitHub Repository][github] as a Pull Request.

## <a name="submit"></a> Submission Guidelines

### Submitting an Issue
Before you submit your issue search the archive, maybe your question was already
answered.

If your issue appears to be a bug, and hasn't been reported, open a new issue.
Help us to maximize the effort we can spend fixing issues and adding new
features, by not reporting duplicate issues.
Providing the following information will increase the chances of your issue
being dealt with quickly:

* **Overview of the Issue** - if an error is being thrown a stack trace helps
* **Motivation for or Use Case** - explain why this is a bug for you
* **Goava Version(s)** - is it a regression?
* **Reproduce the Error** - provide a live example or a unambiguous set of steps
* **Related Issues** - has a similar issue been reported before?
* **Suggest a Fix** - if you can't fix the bug yourself, perhaps you can point
to what might be causing the problem (line of code or commit)

**If you get help, help others. Good karma rulez!**

Here's a template to get you started:

```
Goava version:

What steps will reproduce the problem:
1.
2.
3.

What is the expected result?

What happens instead of that?

Please provide any other information below.
```

### Submitting a Pull Request
Before you submit your pull request consider the following guidelines:

* Search [GitHub](https://github.com/abc-inc/goava/pulls) for an open or closed
  Pull Request that relates to your submission.
  You don't want to duplicate effort.

* Make your changes in a new git branch:
     ```shell script
     git checkout -b my-fix-branch master
     ```

* Create your patch, **including appropriate test cases**.
  The project already has good test coverage, so look at some existing tests if
  you're unsure how to go about it.

* All contributions must be licensed Apache 2.0 and all files must have a copy
  of the boilerplate license comment (can be copied from an existing file).

* Go files should be formatted according to [gofmt][].

* Please squash all commits for a change into a single commit (this can be done
  using `git rebase -i`). Do your best to have a [well-formed commit message][]
  for the change.

* Run the full Goava test suite, and ensure that all tests pass.

* Avoid checking in files that shouldn't be tracked (e.g `.tmp`, `.idea`).
  We recommend using a [global .gitignore][] for this.

* Commit your changes using a descriptive commit message.
     ```shell script
     git commit -a -m <descriptive message>
     ```
  Note: the optional commit `-a` command line option will automatically "add"
  and "rm" edited files.

* Push your branch to GitHub:
    ```shell script
    git push origin my-fix-branch
    ```

* In GitHub, send a pull request to `Goava:master`.

* If we suggest changes then:
  * Make the required updates.
  * Re-run the Goava test suite to ensure tests are still passing.
  * Rebase your branch and force push to your repository
   (this will update your Pull Request):
    ```shell script
    git rebase master -i
    git push origin my-fix-branch -f
    ```

That's it! Thank you for your contribution!

[github]: https://github.com/abc-inc/goava
[issue tracker]: https://github.com/abc-inc/goava/issues
[gofmt]: https://golang.org/cmd/gofmt/
[well-formed commit message]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
[global .gitignore]: https://help.github.com/articles/ignoring-files/#create-a-global-gitignore

#### After your pull request is merged

After your pull request is merged, you can safely delete your branch and pull
the changes from the main (upstream) repository:

* Delete the remote branch on GitHub either through the GitHub web UI or your
local shell as follows:
    ```shell script
    git push origin --delete my-fix-branch
    ```

* Check out the master branch:
    ```shell script
    git checkout master -f
    ```

* Delete the local branch:
    ```shell script
    git branch -D my-fix-branch
    ```

* Update your master with the latest upstream version:
    ```shell script
    git pull --ff upstream master
    ```
