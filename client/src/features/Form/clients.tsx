import Submit from "features/Reused/Submit";
import { FormContainer, TextFieldElement } from "react-hook-form-mui";

function ClientForm() {
  const handleSubmit = (data: any) => {
    console.log(data);
  };
  return (
    <FormContainer
      //   defaultValues={{
      //     name: "",
      //   }}
      onSuccess={handleSubmit}
    >
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="lname"
        label="Last Name"
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="fname"
        label="First Name"
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="mname"
        label="Middle Name"
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name={"email"}
        label={"Email"}
        required
        type={"email"}
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name={"phone"}
        label={"Phone"}
        required
        type={"tel"}
      />
      <Submit />
    </FormContainer>
  );
}

export default ClientForm;
