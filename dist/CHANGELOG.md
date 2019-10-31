## Changelog

a251053 Add goreleaser
8ef0d4b Remove deadcode variable and test cases sorted
1368173 Unexport followingUser struct because it does not want to be exported
1345788 Maintain consistency of printing followers
c68eb4a Add mock generator command in make file
bc2dd9e Fix test case that broke build
d7e893f Merge pull request #7 from raf924/feat(following)
bb27ac0 Remove info_test.php since it contain unwanted test cases
5762ab8 refactor
3b5a474 cli => cmd fix
899c557 added tests + refactor
1b9c7fb Added command to retrieve following users
5790b2b Add more test coverage for makeRequest()
9892055 Merge pull request #11 from thamaraiselvam/list_of_followers
fae8f72 Make sure that followers empty when response is not a valid json
e41ef7b Merge branch 'master' into list_of_followers
b9bc0ae Refactor info command
0e20e7b Add new command to list followers
8b78e0c Add more unit tests
d52bdb8 Rename cli to cmd
72a68dd Merge pull request #9 from thamaraiselvam/add_list_command
cee898b Add List command
8fd3d70 Refactor service for making request
618e7d4 Remove unwanted block from readme file
86aabb1 Build badge error fixed
1ddaacd added sheild.io badges for travis build status and codecov
7e0f268 added codecov in travis.yml
e46e563 moved httpConfig creation to root.go for reusability
8aa24c0 moved  adding subcommends into root.go
4c31e74 added unit test for invalid username for info command
c5926c1 renamed api to service
c7a5304 added deepsource file
73396d9 changed application content in readme file
5e87221 renamed cli name to gac
1046d90 renamed user command into info
87054da Update README.md
df1cd36 added Makefile with commands fmt vet lint compile
e75f395 fixed linter issues using fmt
20fd564 removed unused binary file git-api-cli
c02dd38 added out dir in .gitignore
0ccfb62 added go lint script
e9c9428 added todo commands in readme file
944e253 Update README.md
0e0aa6f Create README.md
885a84d Initial commit, get User info
