"use client";

import * as React from "react";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Divider from "@mui/material/Divider";
import Input from "@mui/material/Input";
import ListItemText from "@mui/material/ListItemText";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import {
  getTasks,
  addTask,
  deleteTask,
  undoRequest,
  UndoAction,
} from "@/api/tasks";
import PendingIcon from "@mui/icons-material/Pending";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import DeleteIcon from "@mui/icons-material/Delete";
import Typography from "@mui/material/Typography";
import { FormGroup, IconButton } from "@mui/material";
import { Button } from "@mui/base";

type TaskStatusIconMap = {
  [key: string]: JSX.Element;
};

const TaskStatusIcon: TaskStatusIconMap = {
  pending: <PendingIcon />,
  completed: <CheckCircleIcon />,
};

function TaskStatus(props: { status: string }): JSX.Element {
  const icon = TaskStatusIcon[props.status];

  if (icon) {
    return icon;
  }

  return <PendingIcon />;
}

type Task = {
  id: string;
  title: string;
  description: string;
  status: string;
};

type TaskRequest = {
  title: string;
  description: string;
};

export default function TasksList() {
  const [addingTask, setAddingTask] = React.useState(false);
  const [tasks, setTasks] = React.useState<Task[]>([]);
  const [newTask, setNewTask] = React.useState<TaskRequest>({
    title: "",
    description: "",
  });
  const [undo, setUndo] = React.useState<UndoAction | null>(null);

  const refreshTasks = () => getTasks().then(setTasks);

  React.useEffect(() => {
    refreshTasks();
  }, []);

  return (
    <>
      <Typography variant="h4" component="h3" gutterBottom>
        Tasks
      </Typography>
      {
        // TODO - Show Undo as a banner
      }
      {undo && (
        <Button
          id="undo"
          onClick={() => {
            undoRequest(undo)
              .then(refreshTasks)
              .then(() => setUndo(null));
          }}
        >
          Undo
        </Button>
      )}
      <List id="task-list">
        {tasks.map((task: Task) => (
          <>
            <ListItem
              id={`task-${task.id}`}
              alignItems="flex-start"
              key={task.id}
              secondaryAction={
                <IconButton
                  edge="end"
                  aria-label="delete-task"
                  onClick={async () => {
                    // FIXME: handle potential error
                    await deleteTask(task.id).then((undo: UndoAction) => {
                      setUndo(undo);
                      refreshTasks();
                    });
                  }}
                >
                  <DeleteIcon />
                </IconButton>
              }
            >
              <ListItemAvatar>
                <TaskStatus status={task.status} />
              </ListItemAvatar>
              <ListItemText primary={task.title} secondary={task.description} />
            </ListItem>
            <Divider variant="inset" component="li" />
          </>
        ))}
      </List>
      {!addingTask && (
        <Button
          color="primary"
          onClick={() => {
            setAddingTask(true);
          }}
        >
          Add Task
        </Button>
      )}
      {
        // Form to add a new task
      }
      {addingTask && (
        <FormGroup>
          <Input
            id="task-title"
            placeholder="Title"
            onChange={(e) =>
              setNewTask((t) => ({ ...t, title: e.target.value }))
            }
          />
          <Input
            id="task-description"
            placeholder="Description"
            onChange={(e) =>
              setNewTask((t) => ({ ...t, description: e.target.value }))
            }
          />
          <Button
            color="primary"
            onClick={() => {
              setAddingTask(false);
              setNewTask({ title: "", description: "" });
            }}
          >
            Cancel
          </Button>
          <Button
            color="primary"
            onClick={async () => {
              // FIXME: handle potential error
              await addTask(newTask.title, newTask.description).then(
                refreshTasks
              );
              setAddingTask(false);
              setNewTask({ title: "", description: "" });
            }}
          >
            Save
          </Button>
        </FormGroup>
      )}
    </>
  );
}
