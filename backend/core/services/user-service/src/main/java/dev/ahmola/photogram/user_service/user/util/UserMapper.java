package dev.ahmola.photogram.user_service.user.util;

import dev.ahmola.photogram.user_service.user.dto.UserRequest;
import dev.ahmola.photogram.user_service.user.dto.UserResponse;
import dev.ahmola.photogram.user_service.user.model.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class UserMapper {

    public static UserResponse userToResponse(User user) {
        return UserResponse.builder()
                .id(user.getId())
                .username(user.getUsername())
                .bio(user.getBio())
                .avatarUrl(user.getAvatarUrl())
                .createdAt(String.valueOf(user.getCreatedAt()))
                .build();
    }

    public static User requestToUser(UserRequest userRequest){
        return User.builder()
                .id(userRequest.id())
                .username(userRequest.username())
                .passwordHash(userRequest.passwordHash())
                .bio(userRequest.bio())
                .avatarUrl(userRequest.avatarUrl())
                .build();
    }
}
