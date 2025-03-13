import { create } from "zustand";
import {
  getTaskById,
  getTasks,
  //   createTask,
  //   updateTask,
  //   deleteTask,
  Task,
} from "../api/task";

interface TaskStore {
  tasks: Task[];
  fetchTasks: () => Promise<void>;
  fetchTask: (taskID: number) => Promise<void>;
}

export const useTaskStore = create<TaskStore>((set) => ({
  tasks: [],

  fetchTasks: async () => {
    try {
      const tasks = await getTasks();
      set({ tasks });
    } catch (error) {
      console.error("Failed to fetch tasks:", error);
    }
  },

  fetchTask: async (taskID) => {
    try {
      const task = await getTaskById(taskID);
      //set({ task });
    } catch (error) {
      console.error("Failed to fetch task:", error);
    }
  },
}));
