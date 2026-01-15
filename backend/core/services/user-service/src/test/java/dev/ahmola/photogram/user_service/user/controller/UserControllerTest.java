package dev.ahmola.photogram.user_service.user.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import dev.ahmola.photogram.user_service.user.dto.UserRequest;
import dev.ahmola.photogram.user_service.user.dto.UserResponse;
import dev.ahmola.photogram.user_service.user.service.UserService;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.BDDMockito.given;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(UserController.class)
class UserControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private UserService userService;

    @Autowired
    private ObjectMapper objectMapper;

    @Test
    @DisplayName("회원 생성 성공 테스트")
    void createUser_success() throws Exception {
        // given
        UserRequest request = UserRequest.builder()
                .username("test1")
                .passwordHash("1234")
                .bio("hello")
                .avatarUrl("http://example.com/avatar/test1.png")
                .build();
        UserResponse response = UserResponse.builder()
                .id(1L)
                .username("test1")
                .bio("hello")
                .avatarUrl("http://example.com/avatar/test1.png")
                .build();

        // stubbing
        given(userService.createUser(any(UserRequest.class))).willReturn(response);

        // when
        mockMvc.perform(post("/api/v1/user")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                // then
                .andDo(print()) // test log
                .andExpect(status().isCreated()) // 201 Created
                .andExpect(jsonPath("$.userId").value(1L))
                .andExpect(jsonPath("$.username").value("test1"))
                .andExpect(jsonPath("$.bio").value("hello"))
                .andExpect(jsonPath("$.avatarUrl").value("http://example.com/avatar/test1.png"))
                .andExpect(jsonPath("$.createdAt").exists());
    }

    @Test
    @DisplayName("UserId로 회원 조회 성공 테스트")
    void readUser_byUserId_success() throws Exception {
        // given
        Long userId = 1L;
        UserResponse response = UserResponse.builder()
                .id(userId)
                .username("test1")
                .bio("hello")
                .avatarUrl("http://example.com/avatar/test1.png")
                .build();

        // stubbing
        given(userService.getUserByUserId(userId)).willReturn(response);

        // when
        mockMvc.perform(get("/api/v1/user")
                        .param("userId", String.valueOf(userId)))
                // then
                .andDo(print()) // test log
                .andExpect(status().isOk()) // 200 OK
                .andExpect(jsonPath("$.username").value("test1"))
                .andExpect(jsonPath("$.bio").value("hello"))
                .andExpect(jsonPath("$.avatar").value("http://example.com/avatar/test1.png"))
                .andExpect(jsonPath("$.createdAt").exists());

    }

    @Test
    @DisplayName("Username으로 회원 조회 성공 테스트")
    void readUser_byUsername_success() throws Exception {
        // given
        String username = "test1";
        UserResponse response = UserResponse.builder()
                .id(1L)
                .username(username)
                .bio("hello")
                .avatarUrl("http://example.com/avatar/test1.png")
                .build();

        // stubbing
        given(userService.getUserByUsername(username)).willReturn(response);

        // when
        mockMvc.perform(get("/api/v1/user")
                        .param("username", username))
                // then
                .andDo(print())
                .andExpect(status().isOk());
    }
}