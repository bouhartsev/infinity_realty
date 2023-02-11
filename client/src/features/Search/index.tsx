import React, { useEffect, useMemo, useState } from "react";
import {
  Backdrop,
  Grid,
  SpeedDial,
  SpeedDialAction,
  SpeedDialIcon,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAppSelector, useAppDispatch } from "../../app/store";
import styles from "./index.module.css";
import { useDebouncedCallback } from "use-debounce";
import { entities } from "app/entities";

import { tstore } from "app/store";
import SearchItem from "./SearchItem";

function Search() {
  // const count = useAppSelector(selectCount);
  // const dispatch = useAppDispatch();

  const navigate = useNavigate();

  const [cards, setCards] = useState(tstore); // dynamic cards
  const [isActionsOpened, setActionsOpened] = useState(false);
  const [filters, setFilters] = useState<string[]>([]);
  const [search, setSearch] = useState<string>("");

  const handleFilters = (
    event: React.MouseEvent<HTMLElement>,
    newFilters: string[]
  ) => {
    setFilters(newFilters);
  };

  useEffect(() => {
    setCards(
      tstore
        .filter((el) => filters.includes(el.type) || !filters.length)
        .filter((el) =>
          Object.values(el).reduce(function (previous, now) {
            return previous || now.indexOf(search) != -1;
          }, false)
        )
    );
  }, [filters, search]);

  const handleSearch = useDebouncedCallback((newSearch: string) => {
    // console.log(newSearch); // make request
    setSearch(newSearch);
  }, 500);

  const actionButtons = useMemo(
    () =>
      Object.keys(entities).map((key: string) => (
        <SpeedDialAction
          key={key}
          icon={entities[key].icon}
          tooltipTitle={entities[key].label}
          onClick={() => navigate("/form/" + key)}
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
          <Grid item key={card.id} xs={12} sm={6} md={4}>
            <SearchItem data={card} />
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
