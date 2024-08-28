go mod edit -module {NEW_MODULE_NAME}
-- rename all imported module
find . -type f -name '*.go' \
  -exec sed -i -e 's,{OLD_MODULE},{NEW_MODULE},g' {} \;