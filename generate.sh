java  -Dmodels -DmodelDocs=false \
  -jar $HOME/programs/swagger-codegen-cli.jar generate \
  -i swagger.yaml \
  -l go \
  -o models \
  -c config.json