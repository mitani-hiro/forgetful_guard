import { apiClient } from "./client";

export interface Task {
  id: number;
  title: string;
  description: string;
  completed: boolean;
}

export const getTasks = async () => {
  const response = await apiClient.get<Task[]>(`/api/tasks`);
  return response.data;
};

export const getTaskById = async (taskID: number): Promise<Task> => {
  const response = await apiClient.get<Task>(`/api/task/${taskID}`);
  return response.data;
};
