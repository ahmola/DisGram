package dev.ahmola.photogram.common_domain.error;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.context.annotation.Import;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@TestPropertySource(properties = {"spring.cloud.config.enabled=false"})
@WebMvcTest(controllers = TestController.class)
@Import(BaseExceptionHandler.class)
class BaseExceptionHandlerTest {

    @Autowired
    MockMvc mockMvc;

    @Test
    @DisplayName("ApiException is Occured")
    void handleApiException() throws Exception {
        mockMvc.perform(get("/test/api-exception"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.message").value("Entity Not Found"));
    }

    @Test
    @DisplayName("General Exception is Occured")
    void handleGeneral() throws Exception {
        mockMvc.perform(get("/test/general-exception"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.message").value("Unexpected Server Error"));
    }
}

@RestController
class TestController{
    @GetMapping("/test/api-exception")
    public void throwApiException(){
        throw new ApiException(ErrorCode.ENTITY_NOT_FOUND);
    }

    @GetMapping("/test/general-exception")
    public void throwGeneralException(){
        throw new RuntimeException("???");
    }
}