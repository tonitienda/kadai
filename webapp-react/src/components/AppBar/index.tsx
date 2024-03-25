import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import LoginMenu from "@/components/LoginMenu";

export default function MenuAppBar() {
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          Kadai
        </Typography>
        <LoginMenu />
      </Toolbar>
    </AppBar>
  );
}
