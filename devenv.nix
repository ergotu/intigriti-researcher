{
  pkgs,
  lib,
  config,
  inputs,
  ...
}: {
  name = "intigriti-researcher";

  # https://devenv.sh/packages/
  packages = with pkgs; [
    gopls
    air
  ];

  languages = {
    go.enable = true;
  };

  # https://devenv.sh/pre-commit-hooks/
  pre-commit.hooks = {
    check-merge-conflicts.enable = true;
    govet = {
      enable = true;
      pass_filenames = false;
    };
    gotest.enable = true;
    golangci-lint = {
      enable = true;
      pass_filenames = false;
    };
  };

  # Make diffs fantastic
  difftastic.enable = true;

  # https://devenv.sh/integrations/dotenv/
  dotenv.enable = true;

  # https://devenv.sh/integrations/codespaces-devcontainer/
  devcontainer.enable = true;
}
