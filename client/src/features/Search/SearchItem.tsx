import { Card, CardContent, CardHeader, Chip, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import { entities } from "app/entities";
import { useMemo } from "react";

type Props = {
  data: { [key: string]: string };
};

function SearchItem(props: Props) {
  const { data } = props;
  // const data_keys: Array<string> = useMemo(() => Object.keys(data), [data]);
  // const;
  const valuesHeader = useMemo(
    () =>
      Object.keys(data)
        .filter((key) => key.indexOf(entities[data.type]?.hFieldSubstr) !== -1)
        .map((key) => data[key]),
    [data]
  );
  const keysBody = useMemo(
    () =>
      Object.keys(data).filter(
        (key) =>
          !!entities[data.type] &&
          key.indexOf(entities[data.type].hFieldSubstr) === -1 &&
          key !== "id"
      ),
    [data]
  );

  return (
    <Card
      component={Link}
      to={`/form/${data.type + "/" + data.id}`}
      sx={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        textDecoration: "none",
      }}
    >
      <CardHeader
        title={`${valuesHeader.join(" ")}`}
        action={
          <Chip
            color={
              !!entities[data.type].color
                ? entities[data.type].color
                : "default"
            }
            label={data.type}
            sx={{ cursor: "pointer" }}
          />
        }
      />
      {/* <Typography gutterBottom variant="h6" component="h2">
                    
                  </Typography> */}
      {/* <Chip color="primary" sx={{ float: "right" }} /> */}
      <CardContent sx={{ flexGrow: 1 }}>
        {keysBody.map((key: string) => (
          <Typography key={data.id + key}>
            <b>{key}: </b>
            {data[key]}
          </Typography>
        ))}
      </CardContent>
    </Card>
  );
}

export default SearchItem;
