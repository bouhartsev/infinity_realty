import React, { useMemo, useState } from "react";
import {
  Backdrop,
  Card,
  CardContent,
  CardHeader,
  Chip,
  Grid,
  SpeedDial,
  SpeedDialAction,
  SpeedDialIcon,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
  Typography,
} from "@mui/material";
import { Link } from "react-router-dom";
import { useAppSelector, useAppDispatch } from "../../app/store";
import styles from "./index.module.css";
import { useDebouncedCallback } from "use-debounce";
import { entities } from "app/entities";

function Search() {
  // const count = useAppSelector(selectCount);
  // const dispatch = useAppDispatch();

  const cards = ["test", "two", "three", "123", "qwe", "qwedf"]; // dynamic cards
  const [isActionsOpened, setActionsOpened] = useState(false);
  const [filters, setFilters] = useState<string[]>([]);

  const handleFilters = (
    event: React.MouseEvent<HTMLElement>,
    newFilters: string[]
  ) => {
    setFilters(newFilters);
  };

  const handleSearch = useDebouncedCallback((newSearch: string) => {
    console.log(newSearch); // make request
  }, 500);

  const actionButtons = useMemo(
    () =>
      Object.keys(entities).map((key: string) => (
        <SpeedDialAction
          key={key}
          icon={entities[key].icon}
          tooltipTitle={entities[key].label}
          // onClick={handleClose}
        />
      )),
    [entities]
  );

  const toggleButtons = useMemo(
    () =>
      Object.keys(entities)
        .reverse()
        .map((key: string) => (
          <ToggleButton key={key} value={key} color={entities[key].color}>
            {entities[key].label}
          </ToggleButton>
        )),
    [entities]
  );

  return (
    <>
      {/* TODO: Add UP button */}
      <TextField
        label="Search"
        id="search-input"
        fullWidth
        onChange={(e) => handleSearch(e.target.value)}
        // sx={{ m: 1, width: '25ch' }}
      />
      <ToggleButtonGroup
        value={filters}
        onChange={handleFilters}
        aria-label="filter entities"
        sx={{ my: 2 }}
      >
        {toggleButtons}
      </ToggleButtonGroup>
      <Grid container spacing={4}>
        {cards.map((card) => (
          <Grid item key={card} xs={12} sm={6} md={4}>
            <Card
              component={Link}
              to={"/form/"}
              sx={{
                height: "100%",
                display: "flex",
                flexDirection: "column",
                textDecoration: "none",
              }}
            >
              <CardHeader
                title="LastName FirstName MiddleName"
                action={<Chip color={"primary"} label={"clients"} sx={{cursor:"pointer"}}/>}
              />
              {/* <Typography gutterBottom variant="h6" component="h2">
                    
                  </Typography> */}
              {/* <Chip color="primary" sx={{ float: "right" }} /> */}
              <CardContent sx={{ flexGrow: 1 }}>
                <Typography>
                  This is a content of the card. You can use this section to
                  provide any information.
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
      <Backdrop open={isActionsOpened} />
      <SpeedDial
        ariaLabel="Add"
        sx={{ position: "fixed", bottom: 16, right: 16 }}
        icon={<SpeedDialIcon />}
        onClose={() => setActionsOpened(false)}
        onOpen={() => setActionsOpened(true)}
        open={isActionsOpened}
      >
        {actionButtons}
      </SpeedDial>
    </>
  );
}

export default Search;
