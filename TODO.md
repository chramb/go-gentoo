## TODOs
### XML
- [ ] add comments/documentation
- [ ] Create
  - [ ] [glsa.xml](xml/glsa.go)
  - [ ] [metadata.xml]()
    - [ ] support freeform text in tags like `<longdescription/>`
    - [ ] add `<pkg/>` and `<cat/>`
    - [ ] support self-closing tags (`<stabilize-allarches/>`)
  - [ ] [mirrors.xml](xml/mirrors.go)
  - [ ] [projects.xml](xml/projects.go)
  - [X] ~~[repositories.xml](xml/repositories.go)~~
  - [ ] [userinfo.xml](xml/userinfo.go)
- [ ] Add tests for all
- [ ] add `validate()` function