package dev.ahmola.photogram.user_service.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
@Builder
@Data
public class UserResponse {
    Long id;
    String username;
    String bio;
    String avatarUrl;
    String createdAt;
}
