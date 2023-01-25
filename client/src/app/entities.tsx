import {
  Apartment,
  BusinessCenter,
  Hail,
  PlaylistAddCheck,
  RequestPage,
  Search,
  SupervisorAccount,
} from "@mui/icons-material";
import { ReactElement } from "react";

type ColorOptions =
  | "standard"
  | "primary"
  | "secondary"
  | "error"
  | "info"
  | "success"
  | "warning";

export const entities: {
  [key: string]: { label: string; color: ColorOptions; icon: ReactElement };
} = {
  clients: {
    label: "Clients",
    color: "primary",
    icon: <SupervisorAccount />,
  },
  agents: {
    label: "Agents",
    color: "standard",
    icon: <Hail />,
  },
  realty: {
    label: "Real estate",
    color: "error",
    icon: <Apartment />,
  },
  offers: {
    label: "Offers",
    color: "warning",
    icon: <BusinessCenter />,
  },
  demands: {
    label: "Demands",
    color: "success",
    icon: <RequestPage />,
  },
  deals: {
    label: "Deals",
    color: "info",
    icon: <PlaylistAddCheck />,
  },
};
