import { Box, Container, Flex, Center } from "@chakra-ui/react";
import teoriImg from "../assets/teori.png";
import howImg from "../assets/how.png";

const HowToUse = () => {
  return (
    <Box
      bgGradient="linear(to-b, white, #e9dcdc, #cbbdbd)"
      minHeight="90vh"
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      
      <img
        src={teoriImg}
        alt="Teori"
        style={{ width: "70%", height: "auto", marginTop: "70px"}}
      />

      <img
        src={howImg}
        alt="How to use"
        style={{ width: "70%", height: "auto", marginTop: "70px", marginBottom: "70px"}}
      />

    </Box>
  );
};

export default HowToUse;
