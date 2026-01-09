package dev.ahmola.photogram.common_domain.event;

import lombok.*;

@Getter@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class PostCreatedEvent {
    private Long postId;
    private String deletedAt;
}
