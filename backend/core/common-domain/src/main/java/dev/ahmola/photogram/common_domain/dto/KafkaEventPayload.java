package dev.ahmola.photogram.common_domain.dto;


import lombok.*;

@Getter @Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class KafkaEventPayload {
    private String topic;
    private Object payload;
}
