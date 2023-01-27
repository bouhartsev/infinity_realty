import { Route, Routes } from "react-router-dom";
import { CssBaseline, Container } from "@mui/material";
import List from 'features/Search';
import Form from 'features/Form';
// import './App.css';

// import Header from "features/Header";

function App() {
  return (
    <>
      <CssBaseline />
      {/* <Header /> */}
      <Container maxWidth="xl" component="main" sx={{ mt: 4 }}>
          <Routes>
            <Route path="/" element={<List />} />
            <Route path="/form/:type" element={<Form />} />

            {/* <Route path="*" element={<ErrorPage code="404" />} /> */}
          </Routes>
      </Container>
    </>
  );
}

export default App;
