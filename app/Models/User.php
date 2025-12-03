<?php

namespace App\Models;

// use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;
use Illuminate\Database\Eloquent\Relations\HasOne;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Database\Eloquent\Relations\HasMany;

class User extends Authenticatable
{
    /** @use HasFactory<\Database\Factories\UserFactory> */
    use HasFactory, HasUuids, Notifiable;

    public $incrementing = false;

    protected $keyType = 'string';

    /**
     * The attributes that are mass assignable.
     *
     * @var list<string>
     */
    protected $fillable = [
        'name',
        'email',
        'first_name',
        'last_name',
        'identification_number',
        'dob',
        'phone',
        'address',
        'password',
        'profile_picture',
        'is_active',
        'role_id',
    ];

    /**
     * The attributes that should be hidden for serialization.
     *
     * @var list<string>
     */
    protected $hidden = [
        'password',
        'remember_token',
    ];

    /**
     * Get the attributes that should be cast.
     *
     * @return array<string, string>
     */
    protected function casts(): array
    {
        return [
            'email_verified_at' => 'datetime',
            'dob' => 'date',
            'address' => 'array',
            'is_active' => 'boolean',
            'password' => 'hashed',
        ];
    }

    /**
     * @return BelongsTo<Role, self>
     */
    public function role(): BelongsTo
    {
        return $this->belongsTo(Role::class);
    }

    /**
     * @return BelongsToMany<Permissions>
     */
    public function permissions(): BelongsToMany
    {
        return $this->belongsToMany(Permissions::class, 'user_permissions', 'user_id', 'permission_id')
            ->withPivot('id')
            ->withTimestamps();
    }

    /**
     * @return HasOne<Student>
     */
    public function student(): HasOne
    {
        return $this->hasOne(Student::class);
    }

    /**
     * @return HasOne<Professor>
     */
    public function professor(): HasOne
    {
        return $this->hasOne(Professor::class);
    }

    /**
     * Domain notifications stored in notifications table, separate from Laravel's polymorphic notifications channel.
     *
     * @return HasMany<Notification>
     */
    public function notificationsRelation(): HasMany
    {
        return $this->hasMany(Notification::class);
    }

    /**
     * @return HasMany<Trace>
     */
    public function traces(): HasMany
    {
        return $this->hasMany(Trace::class);
    }

    /**
     * @return HasMany<UserInvitation>
     */
    public function invitations(): HasMany
    {
        return $this->hasMany(UserInvitation::class);
    }
}
