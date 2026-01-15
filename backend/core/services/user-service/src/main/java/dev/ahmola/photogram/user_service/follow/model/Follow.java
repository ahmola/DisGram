package dev.ahmola.photogram.user_service.follow.model;

import dev.ahmola.photogram.common_domain.domain.BaseEntity;
import dev.ahmola.photogram.user_service.user.model.User;
import jakarta.persistence.*;
import lombok.*;

@Getter@Setter@ToString
@AllArgsConstructor
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
