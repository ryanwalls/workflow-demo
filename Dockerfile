FROM microsoft/aspnet:1.0.0-rc1-update1-coreclr

COPY . /app
WORKDIR /app
RUN ["dnu", "restore"]
RUN ["dnx", "-p", "test/Workflow.UnitTests/project.json", "test"]

ENTRYPOINT ["dnx", "-p", "src/Workflow/project.json", "run"]
