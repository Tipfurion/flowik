type Flow = {
  id: string;
  type: "flow";
  name?: string;
  description?: string;
  steps: (Flow | Job)[];
};

type Job = {
  id: string;
  type: "job";
  name?: string;
  image: string;
  command?: "cmd";
  args?: Record<string, string>;
  timeout?: number;
  environment?: Record<string, string>;
};
