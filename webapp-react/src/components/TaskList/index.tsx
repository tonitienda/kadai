"use client";

import * as React from "react";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Divider from "@mui/material/Divider";
import ListItemText from "@mui/material/ListItemText";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import { getTasks } from "@/api/tasks";
import PendingIcon from "@mui/icons-material/Pending";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";

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
  ID: string;
  title: string;
  description: string;
  status: string;
};

export default async function TasksList() {
  let tasks = [];
  try {
    tasks = await getTasks();
  } catch (error) {
    console.error(error);
    return null;
  }

  return (
    <>
      <List>
        {tasks.map((task: Task) => (
          <>
            <ListItem alignItems="flex-start" key={task.ID}>
              <ListItemAvatar>
                <TaskStatus status={task.status} />
              </ListItemAvatar>
              <ListItemText primary={task.title} secondary={task.description} />
            </ListItem>
            <Divider variant="inset" component="li" />
          </>
        ))}
      </List>
    </>
  );
}
