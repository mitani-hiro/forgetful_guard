import { apiClient } from "./client";
import { components } from "./openapi.gen";

type Task = components["schemas"]["Task"];

export const getTasks = async () => {
  const response = await apiClient.get<Task[]>(`/api/tasks`);
  return response.data;
};

export const getTaskById = async (taskID: number): Promise<Task> => {
  const response = await apiClient.get<Task>(`/api/task/${taskID}`);
  return response.data;
};
