package dev.ahmola.photogram.user_service.follow.dto;

import lombok.Builder;

@Builder
public record FollowResponse(
        Long followerId,
        Long followeeId,
        String createdAt) {}
