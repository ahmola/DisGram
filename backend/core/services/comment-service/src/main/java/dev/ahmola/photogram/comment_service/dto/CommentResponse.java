package dev.ahmola.photogram.comment_service.dto;

import lombok.Builder;

@Builder
public record CommentResponse(
    Long id,
    Long userId,
    Long postId,
    String content,
    String createdAt
) { }
