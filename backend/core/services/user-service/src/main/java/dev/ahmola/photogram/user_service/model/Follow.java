package dev.ahmola.photogram.user_service.model;

import dev.ahmola.photogram.common_domain.domain.BaseEntity;
import jakarta.persistence.*;
import lombok.*;

@Getter@Setter@ToString
@NoArgsConstructor
@Builder
@Entity
@Table(name = "follows")
@IdClass(FollowId.class)
public class Follow extends BaseEntity {

    @Id
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(
            name = "follower_id",
            nullable = false,
            foreignKey = @ForeignKey(name = "fk_follows_follower")
    )
    private User follower;

    @Id
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(
            name = "followee_id",
            nullable = false,
            foreignKey = @ForeignKey(name = "fk_follows_followee")
    )
    private User followee;
}
