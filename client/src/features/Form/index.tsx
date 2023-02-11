import { Container } from "@mui/material";
import React from "react";
import { FC } from "react";
import { useParams } from "react-router-dom";

import clients from "./clients";

function FormContainer() {
  const { type } = useParams();
  const forms: { [key: string]: FC } = {
    clients,
    //   Agents,
  };
  return (
    <>
      {/* Add header with tabs and then swipeable views */}
      <Container maxWidth="xs" sx={{ mx: "auto" }}>
        {!!type && type in forms && React.createElement(forms[type])}
      </Container>
    </>
  );
}

export default FormContainer;
