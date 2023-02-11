import {
  Apartment,
  BusinessCenter,
  Hail,
  PlaylistAddCheck,
  RequestPage,
  SupervisorAccount,
} from "@mui/icons-material";
import { ReactElement } from "react";

type ColorOptions =
  | "primary"
  | "secondary"
  | "error"
  | "info"
  | "success"
  | "warning";

export const entities: {
  [key: string]: {
    label: string; // official label
    color?: ColorOptions; // color for menus, filters and badges
    icon: ReactElement; // icon for menus, filters and pages
    hFieldSubstr: string; // Substring to choose fields for card's header from items
  };
} = {
  clients: {
    label: "Clients",
    color: "primary",
    icon: <SupervisorAccount />,
    hFieldSubstr: "name",
  },
  agents: {
    label: "Agents",
    icon: <Hail />,
    hFieldSubstr: "name",
  },
  realty: {
    label: "Real estate",
    color: "error",
    icon: <Apartment />,
    hFieldSubstr: "coord",
  },
  offers: {
    label: "Offers",
    color: "warning",
    icon: <BusinessCenter />,
    hFieldSubstr: "realty",
  },
  demands: {
    label: "Demands",
    color: "success",
    icon: <RequestPage />,
    hFieldSubstr: "realty",
  },
  deals: {
    label: "Deals",
    color: "info",
    icon: <PlaylistAddCheck />,
    hFieldSubstr: "realty",
  },
};
