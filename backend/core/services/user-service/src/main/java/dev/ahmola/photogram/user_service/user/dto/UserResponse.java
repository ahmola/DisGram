package dev.ahmola.photogram.user_service.user.dto;

import lombok.Builder;

@Builder
public record UserResponse(
        Long id,
        String username,
        String bio,
        String avatarUrl,
        String createdAt) {}