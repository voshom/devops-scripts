// types.ts
export enum ProjectType {
  Frontend = 'frontend',
  Backend = 'backend',
  Fullstack = 'fullstack',
}

export enum DeploymentTarget {
  Local = 'local',
  Prod = 'prod',
  Dev = 'dev',
}

export interface ProjectConfig {
  name: string;
  type: ProjectType;
  deploymentTarget: DeploymentTarget;
}

export interface EnvironmentVariables {
  [key: string]: string;
}

export interface Config {
  project: ProjectConfig;
  env: EnvironmentVariables;
}