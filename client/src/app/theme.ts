import { createTheme } from "@mui/material/styles";

// all colors from brandbook
// some colors (not mentioned) are inherited from lib - https://mui.com/material-ui/customization/palette/#default-values
// * "not bb" - not from brandbook

const theme = createTheme({
  palette: {
    primary: {
      light: "#04A0FF",
      main: "#0091EA",
      dark: "#0081D1",
    },
    secondary: {
      light: "#FFF6A5",
      main: "#FFEE68",
      dark: "#FFE60C",
    },
    warning: {
      light: "#FF9C1A",
      main: "#FF9100",
      dark: "#E68300",
    },
    error: {
      light: "#FF4A6D",
      main: "#FF1744",
      dark: "#E3002C",
    },
    success: {
      light: "#00D9BB",
      main: "#00BFA5",
      dark: "#00A68F",
    },
    text: {
      primary: "#37474F",
      secondary: "#78909C",
      disabled: "#B0BEC5",
    },
  },
  // components: {
  //     MuiCssBaseline: {
  //         styleOverrides: require("assets/fonts.css"),
  //     },
  // },
});

export const tableBaseSX = {
  "&.MuiDataGrid-root .MuiDataGrid-columnHeader:focus-within, &.MuiDataGrid-root .MuiDataGrid-cell:focus-within":
    {
      outline: "none !important",
    },
  "& .row-status--false": {
    color: "error.contrastText",
  },
  "& .row-status--false::before": {
    bgcolor: "error.light",
  },
};

export default theme;
