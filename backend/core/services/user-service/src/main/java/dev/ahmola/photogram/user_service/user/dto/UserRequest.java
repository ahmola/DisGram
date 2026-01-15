package dev.ahmola.photogram.user_service.user.dto;

import lombok.Builder;

@Builder
public record UserRequest(
        Long id,
        String username,
        String passwordHash,
        String bio,
        String avatarUrl
){}